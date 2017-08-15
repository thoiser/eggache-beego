<html>

<head>

  {{template "inc/meta.tpl" .}}
  <title>视频</title>
</head>

<body>
  {{template "inc/comment.tpl" .}}

  <div class="tm-video-img-container">
    <img src="static/img/demo3.jpg" alt="Image" class="hidden-lg-up img-fluid">
  </div>
  <section class="tm-section">
    <div class="container-fluid">
      <div class="row">
        <div class="col-xs-12 col-sm-12 col-md-12 col-lg-12">
          <div class="list-group">
            {{range $k,$v := .videoList}}
            <a href="/video/{{$v.Id}}" class="list-group-item">
              <h4 class="list-group-item-heading">{{$v.Title}}</h4>
              <div style="margin:10px 0;">
                {{date $v.Ctime "Y-m-d H:i:s"}}
              </div>
            </a>
            <br/>{{end}}
          </div>


          {{template "inc/page.tpl" .}}

        </div>
      </div>
    </div>
  </section>

  {{template "inc/footer.tpl" .}}


</body>

</html>
