import logging
import config
import utility
from os import path
from os import makedirs
from sys import stdout


class log(object):

    def __init__(self):
        self.config = config.config("config/config.ini")
        self.loglevel = self.getLevel(self.config.getValue("LOG", "level"))

        # create logger with easywall
        root = logging.getLogger()
        root.handlers.clear()  # workaround for default stdout handler
        root.setLevel(self.loglevel)

        # create formatter and add it to the handlers
        formatter = logging.Formatter(
            '[%(asctime)s] [%(levelname)s] [%(filename)s:%(lineno)d] %(message)s')

        # create console handler -> logs are always written to stdout
        stdHandler = logging.StreamHandler(stdout)
        stdHandler.setLevel(self.loglevel)
        stdHandler.setFormatter(formatter)
        root.addHandler(stdHandler)

        # create file handler if enabled in configuration
        if bool(self.config.getValue("LOG", "to_files")) == True:
            # create log filepath if not exists
            utility.create_folder_if_not_exists(
                self.config.getValue("LOG", "filepath"))

            fileHandler = logging.FileHandler(self.config.getValue(
                "LOG", "filepath") + "/" + self.config.getValue("LOG", "filename"))
            fileHandler.setLevel(self.loglevel)
            fileHandler.setFormatter(formatter)
            root.addHandler(fileHandler)

    def closeLogging(self):
        root = logging.getLogger()
        for handler in root.handlers:
            handler.close()
            root.removeFilter(handler)

    def getLevel(self, logLevel):
        level = logging.NOTSET
        if logLevel == "debug":
            level = logging.DEBUG
        elif logLevel == "info":
            level = logging.INFO
        elif logLevel == "warning":
            level = logging.WARNING
        elif logLevel == "error":
            level = logging.ERROR
        elif logLevel == "critical":
            level = logging.CRITICAL
        return level