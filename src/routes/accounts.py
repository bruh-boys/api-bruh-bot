import random
from dotenv import load_dotenv
import os

load_dotenv()
email1 = os.getenv('EMAIL1')
email2 = os.getenv('EMAIL2')
password = os.getenv('PASSWORD')

accounts = [email1,email2]  # email accunts

p = password  # password accounts


def selectRandom(accounts):
    return random.choice(accounts)
