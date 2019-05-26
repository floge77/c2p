# cloud2podcast
A Go-Tool for creating a Podcast readable XML/RSS Feed from Youtube/Mixcloud/Soundcloud.

## Purpose
I personally listen a lot to Youtube, Mixcloud and Soundcloud sets by any artists. 

Living in Germany means having the worst mobile data contracts of all developed countries. So streaming of Youtube/Mixcloud/Soundcloud is really not an option to enjoy your favorite music.
I could waste my time every day using one out of many X-cloud downloaders to save my favorite sets and sync them with my mobile phone - that SUCKS!

For me the perfect solution would be to download my favorite sets like my podcasts get downloaded - without any action from my side.
To do so I built this little tool which reads your favorite artists from a yaml file and creates out of that information a XML/RSS Feed which any podcast application can read and download.

## Download Tracks
Use youtube-dl to download tracks. 
I use my raspberry-pi as a podcastprovider and download the channels with the following command:
`youtube-dl --verbose -x -i --dateafter now-12months --audio-format "mp3" --embed-thumbnail --add-metadata --match-filter "duration>1800" --download-archive ./youtube/Q-Dance/archive.txt -o "./youtube/Q-Dance/%(title)s__%(uploader)s__%(upload_date)s.%(ext)s" https://www.youtube.com/user/Qdancedotnl/videos`
Be aware of the output format. Cloud2Podcast expects exactly this outputformat to work.
In the config.yaml which has to be placed in your download directory you have to configure the podcast (see the config.yaml as an example). 
The `channelName` in the config.yaml has to be the directory name in the downloads/youtube directory.
Sounds complicated but it is easy.
Example file tree:
```
downloads
└── youtube
|    └── Q-dance
|       └──file.mp3
└── config.yaml
```

## Installation
* Ensure you've installed [docker](https://docs.docker.com/install/) 
* `git clone https://github.com/floge77/cloud2podcast`
* `cd cloud2podcast && buildCloud2podcast.sh`
* Execute `./runContainer.sh` which starts the container with the downloads directory from your current path (pwd) mounted. The directory structure mentioned above is required.
* Check if the application is running at `http://localhost:8080/<your configured Podcast channel>` (in my example `http://localhost:8080/Q-Dance`)



## Current Status
At the moment cloud2podcast works for youtube channels. This is possible with the usage of the awesome [youtube-dl project](https://github.com/rg3/youtube-dl/).
The configured channel will be downloaded and the information of the files will be transformed to a podcast-feed.

The next big goal is get Mixcloud to work!  


## Mixcloud Examples - Help needed

### Example 1
Artist: "q-dance" Track: "q-dance-presents-hardstyle-top-40-l-october-2018" :
https://www.mixcloud.com/Q-dance/q-dance-presents-hardstyle-top-40-l-october-2018/

Final DL URL (generated via http://www.mixcloud-downloader.com/ ):
http://stream8.mixcloud.com/secure/c/m4a/64/6/7/d/5/9ceb-6bcf-4277-a802-4e3254677cc8.m4a?sig=k0rJaou07CtJCbLw-pT59g

After clicking play on the original URL I can see in the chrome developer tools under network (check at media) following URL:
https://audiocdn6.mixcloud.com/previews/6/7/d/5/9ceb-6bcf-4277-a802-4e3254677cc8.mp3

### Example 2
Artist: "q-dance" Track: "q-dance-presents-next-episode-226-by-uncaged" :
https://www.mixcloud.com/Q-dance/q-dance-presents-next-episode-226-by-uncaged/

Final DL URL (generated via http://www.mixcloud-downloader.com/ ): 
http://stream6.mixcloud.com/secure/c/m4a/64/c/a/6/8/e386-fdc0-4b40-942e-6b53c21d109b.m4a?sig=pyluBBOLIGfsyr8TAE7MdA

After clicking play on the original URL I can see in the chrome developer tools under network (check at media) following URL: https://audiocdn2.mixcloud.com/previews/c/a/6/8/e386-fdc0-4b40-942e-6b53c21d109b.mp3


## Lessons learned (until now)
Seems like we can extract the trackID like: .../previews/ + **trackID** + .mp3

What is missing now is the signature.

Lessons learned:
DL URL: 
http://stream + **an Integer** + .mixcloud.com/secure/c/m4a/64 + **trackID** + **signature**

## Useful links
Python Implementation for downloading tracks from Mixcloud: https://github.com/rg3/youtube-dl/blob/master/youtube_dl/extractor/mixcloud.py

Mixcloud-downloader to verify generated download URLs:
http://www.mixcloud-downloader.com/