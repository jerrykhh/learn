from bs4 import BeautifulSoup
import requests
from datetime import datetime
from selenium import webdriver


class Page :

    def set_id(self, id):
        self.id = id[2: len(id)-2]
    
    def set_url(self, url):
        self.url = url[2: len(url)-2]

    def __str__(self):
        return f"{self.id} {self.url}"

def check_str_index(check_str: str):

    cache_char:str = ""
    local_page_index: int = 0
    count_page_create_index: int = 0
    page_create: bool = False
    pages = []
    temp_page = Page()

    for char in list(check_str):
        
        if "loadPageToDiv" in cache_char:
            page_create = True
            cache_char = ""

        if page_create: 
            if "," in cache_char:
                if count_page_create_index == 0: 
                    temp_page.set_id(cache_char)
                    cache_char = ""
                    count_page_create_index = 1
                elif count_page_create_index == 1:
                    temp_page.set_url(cache_char)
                    cache_char = ""
                    count_page_create_index = 0
                    page_create = False
                    pages.append(temp_page)
                    temp_page = Page()

        cache_char += char
        
    return pages





def main() :
    today = datetime.now()
    html_text = requests.get('https://olympics.com/tokyo-2020/olympic-games/en/results/all-sports/olympic-schedule-and-results-date=2021-07-25.htm').text
    html_str = BeautifulSoup(html_text, "html.parser").prettify().strip()


    pages = check_str_index(html_str)
    chrome = webdriver.Chrome('./driver/chromedriver')
    for page in pages[1:len(pages)-2]:
        
        chrome.get(f"https://olympics.com/tokyo-2020/olympic-games/{page.url[8:len(page.url)]}")

        table = chrome.find_element_by_tag_name("table")
        start_row = 1
        index = 0
        for row in table.find_elements_by_tag_name("tr"):

            try:
                try:
                    if index >= start_row:
                        tds = row.find_elements_by_tag_name("td")
                        time = tds[0].find_element_by_xpath("strong/span").get_attribute("local-time")
                        location = tds[1].text
                        event = tds[2].find_element_by_xpath("div/div/div/div/div/a").text + " "
                        event += tds[2].find_elements_by_xpath("div/div//div/div[@class='d-flex']/div")[1].text
                        try:
                            span = tds[2].find_elements_by_xpath("div/div/div")[1].find_element_by_xpath('div/div/div/div[@class="playerTag"]/span')
                            country = span.find_element_by_xpath("img").get_attribute('title')+ " " + span.find_element_by_xpath("abbr").text
                        except Exception:
                            country = "TBD"
                        try:
                            col_row = tds[2].find_elements_by_xpath("div/div/div/div/div/div")[1].find_element_by_xpath('div[@class="playerTag"]/span')
                            vs_contry = col_row.find_element_by_xpath("img").get_attribute('title') + " " + col_row.find_element_by_xpath("abbr").text
                        except Exception:
                            vs_contry = "TBD"
                        
                        if "HKG" in country or "HKG" in vs_contry:
                            print(f"{time}\n{location}\n{event}\n{country}\n{vs_contry}")
                            print()

                except Exception:
                    th = row.find_element_by_tag_name("th")
                    span = th.find_elements_by_tag_name("span")
                    # print(th.text[0: len(th.text)-3], end=" ")
                    # print(span[0].get_attribute('local-time') + " - " + span[1].get_attribute('local-time'))
            except Exception as e:
                pass

                
            index+=1
            

        # page_html_text = requests.get(f"https://olympics.com/tokyo-2020/olympic-games/{page.url[8:len(page.url)]}").text
        # print(f"https://olympics.com/tokyo-2020/olympic-games/{page.url[8:len(page.url)]}")
        # soup = BeautifulSoup(page_html_text, "html.parser")
        # table = soup.find("table", class_="table-schedule")
        # rows = table.find_all('thead')
        # print(rows)
        # return
        # rows = table_body.find_all('tr')

        # for row in rows:
        #     print(row)
        #     return

        




if __name__ == '__main__':
    main()
