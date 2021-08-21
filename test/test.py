import requests

"""
here im just making a simple request for test

"""

r = requests.post("http://localhost:5000/sms",
                  json={"phone": "+5759075894789"})
print(r.text)
