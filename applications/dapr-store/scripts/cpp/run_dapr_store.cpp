#include <cstdio>
#include <iostream>
#include <memory>
#include <stdexcept>
#include <string>
#include <array>
#include <thread>
#include <ctime>
#include <unistd.h>
#include <mutex>
#include "json.hpp"

using namespace std;

using json = nlohmann::json;

std::recursive_mutex mtx;

string exec(const char* cmd) {
    array<char, 128> buffer;
    string result;
    unique_ptr<FILE, decltype(&pclose)> pipe(popen(cmd, "r"), pclose);
    if (!pipe) {
        throw std::runtime_error("popen() failed!");
    }
    while (fgets(buffer.data(), buffer.size(), pipe.get()) != nullptr) {
        result += buffer.data();
    }
    return result;
}

int get_count(string item, int session_id) {
    std::lock_guard<std::recursive_mutex> lk(mtx);
    string command = "dapr invoke --app-id cart --method get/dip";
//    string command = "dapr invoke --app-id cart --method get/dip_cart_"
//                      + std::to_string(session_id);
    string result = exec(command.c_str());
    string cart = result.substr(0, result.find('\n'));
    auto cart_json = json::parse(cart);
    if (cart_json["products"].find(item) != cart_json["products"].end()) {
        return cart_json["products"][item];
    }
    return 0;
}


void add_item(string item, int session_id) {
    std::lock_guard<std::recursive_mutex> lk(mtx);
    int quantity = get_count(item, session_id);
    string command = "dapr invoke --app-id cart --method setProduct/dip/"
//    string command = "dapr invoke --app-id cart --method setProduct/dip_cart_"
//                      + std::to_string(session_id) + "/"
                      + item + "/" + to_string(quantity + 1);
    string result = exec(command.c_str());
}


void delete_item(string item, int session_id) {
    std::lock_guard<std::recursive_mutex> lk(mtx);
    get_count(item, session_id);  // Read and then delete
    string command = "dapr invoke --app-id cart --method setProduct/dip/"
//    string command = "dapr invoke --app-id cart --method setProduct/dip_cart_"
//                      + std::to_string(session_id) + "/"
                      + item + "/0";
    string result = exec(command.c_str());
}


void clear_cart() {
    std::lock_guard<std::recursive_mutex> lk(mtx);
    string command = "dapr invoke --app-id cart --method clear/dip";
//    string command = "dapr invoke --app-id cart --method clear/dip_cart_1";
    string result = exec(command.c_str());
}


void concurrent_delete(int t_id, string item, int iter_count) {
    if (t_id == 1)
        add_item(item, t_id);
    else
        delete_item(item, t_id);

    int q1 = get_count(item, t_id);
    int q2 = get_count(item, t_id);
    cout << "Iter "<< iter_count << " Thread " << t_id << " Read: "<< q1 << " " <<  q2 << endl;

    if (q1 == 0 && q2 == 2)
        cout << "FOUND VIOLATION at iteration " << iter_count << endl;
}


void checkout_order(int session_id) {
}


void concurrent_checkout(int t_id, string item, int iter_count) {
    if (t_id == 1)
        checkout_order(t_id);
    else
        add_item(item, t_id);

    int q1 = get_count(item, t_id);
    int q2 = get_count(item, t_id);
    cout << "Iter "<< iter_count << " Thread " << t_id << " Read: "<< q1 << " " <<  q2 << endl;

    if (q1 == 0 && q2 == 2)
        cout << "FOUND VIOLATION at iteration " << iter_count << endl;
}


string gen_random(const int len, long seed) {
    string tmp_s;
    static const char alphanum[] =
        "0123456789"
        "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
        "abcdefghijklmnopqrstuvwxyz";

    srand(seed);

    tmp_s.reserve(len);

    for (int i = 0; i < len; ++i)
        tmp_s += alphanum[rand() % (sizeof(alphanum) - 1)];

    return tmp_s;
}


int main() {
    int num_iters = 1000;
    int num_threads = 2;

    cout << exec("dapr invoke --app-id users --method register --payload \'{\"username\":\"dip\", \"displayName\":\"Diptanshu\", \"profileImage\":\"abc.jpg\"}\'");

    string item = "abc";

    for (int i = 0; i < num_iters; i++) {
        vector<thread> threads;
        clear_cart();
        add_item(item, 1);

        for (int j = 0; j < num_threads; j++) {
            threads.push_back(thread(concurrent_delete, j + 1, item, i));
        }

        for (int j = 0; j < num_threads; j++) {
            threads[j].join();
        }
    }
}
