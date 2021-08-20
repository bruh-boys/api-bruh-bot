import requests

"""
here im just making a simple request for test

"""

r = requests.post("http://localhost:5000/sms",
                  json={"phone": "+455345345"})
print(r.text)
