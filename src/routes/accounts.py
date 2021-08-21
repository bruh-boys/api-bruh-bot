import random
from dotenv import load_dotenv
import os

load_dotenv()
email1 = os.getenv('EMAIL1')
email2 = os.getenv('EMAIL2')
email3 = os.getenv('EMAIL3')
email4 = os.getenv('EMAIL4')
email5 = os.getenv('EMAIL5')
password = os.getenv('PASSWORD')

accounts = [email1,email2,email3,email4,email5]  # email accunts

p = password  # password accounts


def selectRandom(accounts):
    return random.choice(accounts)
