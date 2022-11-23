import datetime
a=datetime.datetime.strptime('2020-06-19 15:52:50', "%Y-%m-%d %H:%M:%S")
b=datetime.datetime.strptime('2020-06-19 16:52:50', "%Y-%m-%d %H:%M:%S")
d=b-a
print(d.total_seconds())
review={"timelinef":"ok"}
if "timeline" in review:
  print("yes")
else:
  print("no")