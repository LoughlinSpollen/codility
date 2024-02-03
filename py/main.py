#!/usr/bin/env python3.7
# coding=utf-8

import logging
import sys
import os
LOGGER = logging.getLogger()

def serve():
    LOGGER.info("Running codality challenges")

    try:
        from challenges import lesson1
        lesson1.run()
    except :
        LOGGER.error("Exception running challenges" + str(sys.exc_info()[0]))


if __name__ == '__main__':
    logging.basicConfig(
        level=logging.DEBUG if os.getenv('DEBUG') == 'True' else logging.INFO,
        format="%(asctime)s [%(threadName)-12.12s] [%(levelname)-5.5s]  %(message)s",
        handlers=[
            logging.StreamHandler(sys.stdout)
        ])
    serve()
