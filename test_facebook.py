import json
from facebook_scraper import get_posts, set_noscript
from facebook_scraper import get_group_info

import time

start = time.time()
ginfo = get_group_info("VenomousAndOtherSnakesOfSiam")
print(ginfo['id'])
#479127595592229

## We need to login Facebook with this account in a browser before running the code.

listposts=[]
#for post in get_posts("nintendo", pages=2):
posts = get_posts(group=444832069872572,cookies='cockies.json',options={"allow_extra_requests":True})

# ?fbid=2017781678415978&set=gm.1293309938135935&idorvanity=294155504718055
count=0

for post in posts:
#   print post on the screen
#    print(json.dumps(post,indent=2, default=str))
    for x in post:
       print(x)
#  print json object to file one object/post per line
    
    with open('fb.json', 'a', encoding='utf-8') as f:
        json.dump(post, f, ensure_ascii=False, default=str)
        f.write("\n")
    # print(post['post_text'])
    # for cmt in post['comments_full']:
    #     print(cmt['commenter_name']," : ", cmt['comment_text'].replace("\n",""))
    #     count += 1 + len(cmt["replies"])
    # print(f'''Comments: {post["comments"]}
    #     Likes: {post["likes"]}
    #     Reactions: {post["reaction_count"]} ({post["reactions"]})
    #     len(comments): {len(post["comments_full"])}
    #     len(comments+replies): {count}''')
end = time.time()
print(end - start)