<!DOCTYPE html>
<html>

<head>
  {{template "inc/meta.tpl" .}}
  <title>视频</title>

  <body>
    <a href="/video" class="nav-link">返回视频列表</a>


    <section class="tm-section">
      <div class="container-fluid">
        <div class="row">
          <div class="embed-responsive embed-responsive-16by9">
            <iframe class="embed-responsive-item" src="{{.video.Url}}"></iframe>
          </div>
        </div>

      </div>
    </section>

    <script src="js/jquery-1.11.3.min.js"></script>

  </body>

</html>
