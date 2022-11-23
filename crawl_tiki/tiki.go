package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Item struct {
	ProductId				int `json: "id"`
	SellerProductId   				int    `json: "seller_product_id"`
	OriginalPrice 			int `json: "original_price"`
	QuantitySold 			 	int `json: "quantity_sold[value]"`
	RatingAverage 			float64  `json: "rating_average"`
	ReviewCount 				int `json:"review_count"`
	UrlPath 						string `json:"url_path"`
}

type Tag struct {
	ID        int       `json:"id"`
	// Name      string    `json:"name"`
	// CreatedAt time.Time `json:"created_at"`
}

func MakeRequest(url string, ch chan<-string) {
  start := time.Now()
  resp, _ := http.Get(url)
  secs := time.Since(start).Seconds()
  body, _ := ioutil.ReadAll(resp.Body)
  ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, len(body), url)
}

func main() {
	// file, _ := ioutil.ReadFile("tiki.json")
	bytes:=[]byte(`
	[
  {
    "id": 169848307,
    "sku": "9915949364321",
    "name": "Kệ Wifi Treo Tường 2 Tầng Không Cần Khoan Để Đầu Thu Kỹ Thuật Số, Remote, Điện Thoại - IG407",
    "url_key": "ke-wifi-treo-tuong-2-tang-khong-can-khoan-de-dau-thu-ky-thuan-so-remote-dien-thoai-ig407-p169848307",
    "url_path": "ke-wifi-treo-tuong-2-tang-khong-can-khoan-de-dau-thu-ky-thuan-so-remote-dien-thoai-ig407-p169848307.html?spid=169848309",
    "type": "",
    "author_name": "",
    "book_cover": null,
    "brand_name": "IGA",
    "short_description": "",
    "price": 62000,
    "list_price": 0,
    "badges": [],
    "badges_new": [
      {
        "code": "delivery_info_badge",
        "placement": "delivery_info",
        "text": "Giao tiết kiệm",
        "type": "delivery_info_badge"
      },
      {
        "code": "official_store",
        "icon": "https://salt.tikicdn.com/ts/upload/5d/4c/f7/0261315e75127c2ff73efd7a1f1ffdf2.png",
        "icon_height": 14,
        "icon_width": 68,
        "placement": "top",
        "type": "icon_badge"
      },
      {
        "code": "freeship_plus",
        "placement": "under_rating",
        "text": "Freeship+",
        "type": "under_rating_text"
      },
      {
        "code": "asa_reward_html_badge",
        "placement": "under_price",
        "text": "Tặng tới 7 ASA (2k ₫)<br/>≈ 3.7% hoàn tiền",
        "text_color": "#808089",
        "type": "asa_reward_html"
      }
    ],
    "discount": 0,
    "discount_rate": 0,
    "rating_average": 4.5,
    "review_count": 265,
    "order_count": 0,
    "favourite_count": 0,
    "thumbnail_url": "https://salt.tikicdn.com/cache/280x280/ts/product/c0/20/26/d18b12c8541c13c0a059761cd6a3159d.png",
    "thumbnail_width": 280,
    "thumbnail_height": 280,
    "freegift_items": [],
    "has_ebook": false,
    "inventory_status": "available",
    "is_visible": false,
    "productset_id": 4771,
    "productset_group_name": "",
    "seller": null,
    "is_flower": false,
    "is_gift_card": false,
    "inventory": null,
    "url_attendant_input_form": "",
    "option_color": [],
    "stock_item": null,
    "salable_type": "",
    "seller_product_id": 169848309,
    "installment_info": null,
    "url_review": "",
    "bundle_deal": false,
    "quantity_sold": { "text": "Đã bán 1451", "value": 1451 },
    "video_url": null,
    "tiki_live": null,
    "original_price": 62000,
    "shippable": true,
    "impression_info": [
      {
        "impression_id": "thanos-product-GBUYhvSgd33edb8B",
        "metadata": {
          "price": 62000,
          "rating_average": 4.5,
          "reviews_count": 265,
          "seller_product_id": 169848309
        }
      },
      {
        "impression_id": "bfe4632f-749d-4eb1-bb93-7f98eaae2625",
        "metadata": {
          "product_id": 4771,
          "service_name": "reco",
          "version": "p_category_mpid_listing_v1_202211210600"
        }
      }
    ],
    "advertisement": null,
    "availability": 1,
    "primary_category_path": "1/2/1883/2150/23570/23588",
    "product_reco_score": 0
  }
	]`)
	var data []Item
	
	_ = json.Unmarshal(bytes, &data)
	fmt.Printf("data : %+v", data[0])
  // start := time.Now()
  // ch := make(chan string)
  // for _,url := range os.Args[1:]{
  //     go MakeRequest(url, ch)
  // }
  // for range os.Args[1:]{
  //   fmt.Println(<-ch)
  // }
  // fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}