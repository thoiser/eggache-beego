<!DOCTYPE html>
<html>

<head>
  {{template "inc/meta.tpl" .}}
  <title>相册</title>
</head>
<form enctype="multipart/form-data" method="post" action="/up" id="form1" style="display:none;">
  <input type="text" name="title" placeholder="输入标题">
  <input type="file" name="file_name" />
  <input type="submit">
</form>

<body>
  {{template "inc/comment.tpl" .}}

  <div class="tm-home-img-container">
    <img src="static/img/demo4.jpg" alt="Image" class="hidden-lg-up img-fluid">
    <!--    <img src="img/tm-home-img.jpg" alt="Image" class="hidden-lg-up img-fluid">-->
  </div>

  <section class="tm-section">
    <div class="container-fluid">
      <div class="row">
        <div class="col-xs-12 col-sm-12 col-md-12 col-lg-12 text-xs-center">
          <h2 class="tm-gold-text tm-title" id="up">Our’s memory</h2>
          <p class="tm-subtitle" id="out">一点一滴的记录</p>
        </div>
      </div>
      <div class="row">

        {{range $k,$v := .photoList}}
        <div class="col-xs-12 col-sm-6 col-md-6 col-lg-3 col-xl-3">

          <div class="tm-content-box">
            <img src="{{$v.Url}}" alt="Image" class="tm-margin-b-20 img-fluid">
            <h4 class="tm-margin-b-20 tm-gold-text">{{$v.Title}}
                            <button class="btn btn-circle btn-danger-outline delete" style="float:right; text-align:right;" type="button"data-id="{{$v.Id}}" >X
                            </button></h4>
          </div>
        </div>{{end}}
      </div>

      {{template "inc/page.tpl" .}}

    </div>
  </section>


  {{template "inc/footer.tpl" .}}
  <script>
    $('#up').click(function() {
      $('#form1').show();
    });
    $('#out').click(function() {
      $('#logout').show();
    });

    //删除楼层
    $('.delete').on('click', function() {
      var t = confirm('您确定要删除吗？');
      if (!t) {
        return false;
      }
      $.ajax({
        url: '/photo/del/' + $(this).attr('data-id'),
        type: 'GET',
        data: {
          id: $(this).attr('data-id')
        },
        success: function(ret) {
          location.href = "/photo";
        }
      });
    });
  </script>

</body>

</html>
