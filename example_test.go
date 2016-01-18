package gimeo_test

func ExampleVimeo() {
	Vimeo("0e6b4801se9725a5261b1eace9co64f353fe0a90", "potato", "accessToken")

	Vimeo("0e6b4801se9725a5261b1eace9co64f353fe0a90", "potato", nil)
}

func ExampleRequest() {
	client := Vimeo("0e6b4801se9725a5261b1eace9co64f353fe0a90", "potato", "accessToken")
	client.Request(`
	  {
	    "path"  : "/me/videos",
	    "query" : {
					"page" : "1"
	        "per_page" : "1"
	  }
	`)
}

func ExampleUploadVideo() {
	client := Vimeo("0e6b4801se9725a5261b1eace9co64f353fe0a90", "potato", "accessToken")
	client.uploadVideo("/path/to/video.mp4", nil)
	client.uploadVideo("/path/to/video.mp4", "/videos/17178")
}

func ExampleGenerateUnauthAccessToken() {
	client := Vimeo("0e6b4801se9725a5261b1eace9co64f353fe0a90", "potato", "accessToken")
	client.GenerateUnauthAccessToken([]string{"create", "public"})
}
