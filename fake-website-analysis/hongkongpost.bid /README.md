### (Scam Web) hongkongpost.bid 分析(Scam Web) hongkongpost.bid 分析
最近屋企人收到一個信息，內容係「郵件因為地址未能送遞」之後引導去一個假嘅郵政局網站(hongkongpost.bid)


咁我就試下分析呢個假網站嘅內容，[pic1]雖然我未試過改收貨地址，但感覺上個網頁style都好似，而且click佢個logo係會redirect去真正嘅郵政局網址，而其他掣都係「裝假狗」都係唔會有反應，包括最上嘅郵遞服務、大量投寄服務、郵政服務、轉語言等等。應該唯一work嘅係你填自己嘅個人資料之後㩒繼續，如果你冇填野會彈個alert出嚟，咁我又睇下佢啲code[pic2]，可以見到佢有兩個function 一個係cn()同hk()，所以初步推論呢個犯罪集團應該係會中國同香港兩邊呃人。

![image](https://github.com/jerrykhh/p-warehouse/blob/master/fake-website-analysis/hongkongpost.bid/images/image1.jpeg)

![image](https://github.com/jerrykhh/p-warehouse/blob/master/fake-website-analysis/hongkongpost.bid/images/image2.jpeg)


之後，我用我嘅方式唔填資料跳落下一頁，可以見到有兩個option，但其實係迫你揀安排新嘅送遞同比錢，呢個多左個work嘅掣係登出，之係㩒提交。[pic3]

![image](https://github.com/jerrykhh/p-warehouse/blob/master/fake-website-analysis/hongkongpost.bid/images/image3.jpeg)


再之後，佢叫你填你嘅信用卡資料，因為心急求其填左之後就㩒比錢，冇cap到圖，當然啦我係填假嘅資料，之後佢網頁就會彈左佢loading嘅畫面[pic4]，之後唔知係咪我用我嘅方法唔填啲資料搞到佢網頁卡左係loading個畫面，唔知佢係咪會扣左$21，當然我都冇呢個勇氣試~

![image](https://github.com/jerrykhh/p-warehouse/blob/master/fake-website-analysis/hongkongpost.bid/images/image4.jpeg)

好，我地開始分析佢嘅程式碼[pic5]，可以見到佢會偷偷地send你啲個人資料同信用卡資料等去一個logs.center嘅網站，唔知佢有冇係頭先嗰頭碌左你張卡，如果係受害人嘅角度可能比左$21姐，唔緊要，但係到見到唔係，你會係唔知幾時嘅時候比人偷碌呢張卡/比人賣去deep web...，所以如果你有冇財物損失都好都一定要cut卡/換卡！

![image](https://github.com/jerrykhh/p-warehouse/blob/master/fake-website-analysis/hongkongpost.bid/images/image5.jpeg)


偷偷地record你資料嘅網站(logs.center)，要登入。[pic6] check下WHOIS database犯罪集團係用namecheap.com嚟reg 呢個domain. [pic7]

![image](https://github.com/jerrykhh/p-warehouse/blob/master/fake-website-analysis/hongkongpost.bid/images/image7.jpeg)

試吓作一個假嘅data send去佢server，data入面有張8964坦克嘅張轉左做base64 code， 之後佢個網有Error可見佢係用laravel php。[pic8]

![image](https://github.com/jerrykhh/p-warehouse/blob/master/fake-website-analysis/hongkongpost.bid/images/image8.jpeg)

Error左就費事debug，直接拎番佢個data做，之後睇佢會return個id， 而我個id係283，咁估計佢record左~280個人嘅data
咁我好唔小心咁寫左啲code係做fake data ，咁我由好唔小心寫左個冇限loop係不斷send request比logs.center呢個網[pic9]

![image](https://github.com/jerrykhh/p-warehouse/blob/master/fake-website-analysis/hongkongpost.bid/images/image9.jpeg)

最後，發現呢個唔小心嘅時候佢return番嘅id係已經由最初283變左1萬幾，犯罪集團就fatfat喇，10000 乘21都有成21萬。[pic10]

![image](https://github.com/jerrykhh/p-warehouse/blob/master/fake-website-analysis/hongkongpost.bid/images/image10.jpeg)

寫呢篇野之前，我已經report左比namecheap.com依家呢個假網站(hongkongpost.bid)已經上唔到了 [pic11]

![image](https://github.com/jerrykhh/p-warehouse/blob/master/fake-website-analysis/hongkongpost.bid/images/image11.jpeg)
