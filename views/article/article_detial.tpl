<!DOCTYPE html>
<html>

<head>
  {{template "inc/meta.tpl" .}}
  <title>日记</title>

  <body>
    {{template "inc/comment.tpl" .}}


    <div class="tm-article-img-container">
      <!--    <img src="img/tm-home-img.jpg" alt="Image" class="hidden-lg-up img-fluid">-->
      <img src="img/demo5.jpg" alt="Image" class="hidden-lg-up img-fluid">
    </div>

    <section class="tm-section">
      <div class="container-fluid">
        <div class="row">

          <div class="col-xs-12 col-sm-12 col-md-8 col-lg-9 col-xl-9">
            <div class="tm-blog-post">
              <h3 class="tm-gold-text">{{.art.Title}}</h3>
              <a class="tm-gold-text pull-lg-right" href="/article/edit/{{.art.Id}}">更改</a>
              <br/>
              <a class="tm-gold-text pull-lg-right delete" href="/article/del/{{.art.Id}}">删除</a>
              <br/> {{.art.Content}}
            </div>

          </div>

        </div>

      </div>
    </section>

    <!-- <script src="js/jquery-1.11.3.min.js"></script>
    <script>
      //删除楼层
      $('.delete').on('click', function() {
        var t = confirm('您确定要删除吗？');
        if (!t) {
          return false;
        }
        $.ajax({
          url: 'delArticle.php',
          type: 'GET',
          data: {
            id: $(this).attr('data-id')
          },
          success: function(ret) {
            location.href = "articleList.php";
          }
        });
      });
    </script> -->
    {{template "inc/footer.tpl" .}}

  </body>

</html>
