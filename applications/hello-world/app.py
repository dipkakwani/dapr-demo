# ------------------------------------------------------------
# Copyright (c) Microsoft Corporation.
# Licensed under the MIT License.
# ------------------------------------------------------------

import os
import time
from requests_futures.sessions import FuturesSession

dapr_port = os.getenv("DAPR_HTTP_PORT", 3500)
dapr_post_url = "http://localhost:{}/v1.0/invoke/nodeapp/method/neworder".format(dapr_port)
dapr_get_url = "http://localhost:{}/v1.0/invoke/nodeapp/method/order".format(dapr_port)


def inject_fault():
    os.system("ccm node1 pause")
    # os.system("ccm node1 stop")
    print("node 1 stopped", flush=True)


def recover_fault():
    os.system("ccm node1 resume")
    # os.system("ccm node1 start")
    print("node 1 started", flush=True)


n = 0
session = FuturesSession()
while True:
    n += 1
    message = {"data": {"orderId": n}}

    try:
        print("POST request", flush=True)
        future_post_response = session.post(dapr_post_url, json=message, timeout=7)
        print("POST sent", flush=True)
        inject_fault()
        post_response = future_post_response.result()
        get_response = session.get(dapr_get_url, timeout=7).result()

        if not post_response.ok:
            print("HTTP POST %d => %s" % (post_response.status_code,
                                     post_response.content.decode("utf-8")), flush=True)
        if not get_response.ok:
            print("HTTP GET %d => %s" % (get_response.status_code,
                                     get_response.content.decode("utf-8")), flush=True)

        if post_response.ok and get_response.ok:
            print(get_response.json(), flush=True)
            val = get_response.json()["orderId"]
            print(val, flush=True)
            if val != n:
                print("Invalid read: " + str(val) + " expected: " + str(n), flush=True)
                exit(-1)

        recover_fault()

    except Exception as e:
        print(e, flush=True)

    time.sleep(1)
