import os

SECRET_KEY = os.getenv('SECRET_KEY', os.urandom(24))
