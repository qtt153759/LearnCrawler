from csv import DictWriter
import datetime
import json
import requests

field_names=["RatingAverage","QuantitySold","ReviewCount","OriginalPrice","ReviewThankCount","ProductUsageTime","ReviewContent","ReviewStar"]
class Item(object):
    def __init__ (self,RatingAverage,QuantitySold,ReviewCount, OriginalPrice,ReviewContent,ReviewThankCount,ProductUsageTime,ReviewStar):
        self.RatingAverage = RatingAverage
        self.QuantitySold = QuantitySold
        self.ReviewCount = ReviewCount
        self.OriginalPrice = OriginalPrice
        self.ReviewThankCount=ReviewThankCount
        self.ProductUsageTime=ProductUsageTime
        self.ReviewContent=ReviewContent
        self.ReviewStar=ReviewStar

f = open('tiki.json')
  
# returns JSON object as 
# a dictionary
datas = json.load(f)
custom_header = {'User-Agent': 'Mozilla/5.0','Content-Type':'application/json; charset=utf-8','content-encoding': 'gzip'}
for data in datas:
    seller_id = requests.get("https://tiki.vn/api/v2/products/"+str(data['id'])+"?platform=web&spid="+str(data['seller_product_id']), headers=custom_header).json()['current_seller']['id']
    print("https://tiki.vn/api/v2/products/"+str(data['id'])+"?platform=web&spid="+str(data['seller_product_id']))
    for star in [1,2,3,4,5]:
        reviews = requests.get("https://tiki.vn/api/v2/reviews?limit=20&include=comments,contribute_info,attribute_vote_summary&sort=stars%7C"+str(star)+"&page=1&spid="+str(data['seller_product_id'])
                                +"&product_id="+str(data['id'])+"&seller_id="+str(seller_id), headers=custom_header).json()['data']
        print("https://tiki.vn/api/v2/reviews?limit=20&include=comments,contribute_info,attribute_vote_summary&sort=stars%7C"+str(star)+"&page=1&spid="+str(data['seller_product_id'])
                                +"&product_id="+str(data['id'])+"&seller_id="+str(seller_id))
        for review in reviews:
            ProductUsageTime=0
            if 'timeline' in review:
                ReviewCreatedDate=datetime.datetime.strptime(review['timeline']['review_created_date'], "%Y-%m-%d %H:%M:%S")
                DeliveryDate=datetime.datetime.strptime(review['timeline']['delivery_date'], "%Y-%m-%d %H:%M:%S")
                ProductUsageTime=(ReviewCreatedDate-DeliveryDate).total_seconds()
            item=Item(RatingAverage=data['rating_average'],QuantitySold=data["quantity_sold"]["value"],ReviewCount=data['review_count']
                    ,OriginalPrice=data['original_price'],ReviewContent=review['content'],ProductUsageTime=ProductUsageTime,ReviewThankCount=review['thank_count'],ReviewStar=review['rating'])
            print(item.__dict__)
            with open('reviewTiki.csv', 'a+', newline='') as write_obj:
        # Create a writer object from csv module
                dict_writer = DictWriter(write_obj, fieldnames=field_names)
        # Add dictionary as wor in the csv
                dict_writer.writerow(item.__dict__)
    # data=response.json()
# print(response.content.decode("utf-8"))
# with open('tikihtml.json', 'a', encoding='utf-8') as f:
#         json.dump(data, f, ensure_ascii=False, default=str)
