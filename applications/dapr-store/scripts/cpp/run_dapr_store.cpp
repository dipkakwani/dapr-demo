#include <cstdio>
#include <iostream>
#include <memory>
#include <stdexcept>
#include <string>
#include <array>
#include <thread>
#include <ctime>
#include <unistd.h>
#include "json.hpp"

using namespace std;

using json = nlohmann::json;

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

int get_count(string user, string item, int session_id) {
    string command = "dapr invoke --app-id cart --method get/" + user;
    string result = exec(command.c_str());
    string cart = result.substr(0, result.find('\n'));
    auto cart_json = json::parse(cart);
    if (cart_json["products"].find(item) != cart_json["products"].end()) {
        return cart_json["products"][item];
    }
    return 0;
}


void add_item(string user, string item, int session_id) {
    string command = "dapr invoke --app-id cart --method addItem/" + user + "/" + item;
    string result = exec(command.c_str());
}


void delete_item(string user, string item, int session_id) {
    string command = "dapr invoke --app-id cart --method deleteItem/" + user + "/" + item;
    string result = exec(command.c_str());
}


void clear_cart(string user) {
    string command = "dapr invoke --app-id cart --method clear/" + user;
    string result = exec(command.c_str());
}


void concurrent_delete(int t_id, string user, string item, int iter_count) {
    if (t_id == 1)
        add_item(user, item, t_id);
    else
        delete_item(user, item, t_id);

    int q1 = get_count(user, item, t_id);
    int q2 = get_count(user, item, t_id);
    cout << "Iter "<< iter_count << " Thread " << t_id << " Read: "<< q1 << " " <<  q2 << endl;

    if (q1 == 0 && q2 == 2)
        cout << "FOUND VIOLATION at iteration " << iter_count << endl;
}


void checkout_order(int session_id) {
}


void concurrent_checkout(int t_id, string user, string item, int iter_count) {
    if (t_id == 1)
        checkout_order(t_id);
    else
        add_item(user, item, t_id);

    int q1 = get_count(user, item, t_id);
    int q2 = get_count(user, item, t_id);
    cout << "Iter "<< iter_count << " Thread " << t_id << " Read: "<< q1 << " " <<  q2 << endl;

    if (q1 == 0 && q2 == 2)
        cout << "FOUND VIOLATION at iteration " << iter_count << endl;
}


string gen_random(const int len, long seed) {
    string tmp_s;
    static const char alphanum[] =
        "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
        "abcdefghijklmnopqrstuvwxyz";

    srand(seed);

    tmp_s.reserve(len);

    for (int i = 0; i < len; ++i)
        tmp_s += alphanum[rand() % (sizeof(alphanum) - 1)];

    return tmp_s;
}


int main() {
    int num_iters = 100;
    int num_threads = 2;

//    cout << exec("dapr invoke --app-id users --method register --payload \'{\"username\":\"dip\", \"displayName\":\"Diptanshu\", \"profileImage\":\"abc.jpg\"}\'");

    string item = "abc";

    for (int i = 0; i < num_iters; i++) {
        string user = gen_random(5, i * 11 + 97);
        vector<thread> threads;
        clear_cart(user);
        add_item(user, item, 1);

        for (int j = 0; j < num_threads; j++) {
            threads.push_back(thread(concurrent_delete, j + 1, user, item, i));
        }

        for (int j = 0; j < num_threads; j++) {
            threads[j].join();
        }
    }
}
