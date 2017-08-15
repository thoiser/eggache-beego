<!DOCTYPE html>
<html>

<head>

  {{template "inc/meta.tpl" .}}
  <title>日记</title>
  <style>
    .time {
      padding: 50px 0;
      float: right;
      margin: 50px 0
    }
  </style>
</head>

<body>
  {{template "inc/comment.tpl" .}}

  <section class="tm-section">
    <div class="container-fluid">
      <div class="row">
        <div class="col-xs-12 col-sm-12 col-md-12 col-lg-12">
          <form action="/article/edit/{{.art.Id}}" method="post">
            <div class="input-group input-group-lg">
              <span class="input-group-addon">日记名称</span>
              <input type="text" class="form-control" name="title" placeholder="请输入日记名称" value="{{.art.Title}}">
            </div>
            <br/>
            <div class="input-group input-group-lg">
              <span class="input-group-addon">日记内容</span>
              <textarea type="text" class="form-control" name="content" placeholder="请输入日记内容" style="height: 500px">{{.art.Content}}</textarea>
            </div>
            <br/>
            <input type="hidden" name="id" value="{{.art.Id}}">
            <button type="submit" class="btn btn-warning btn-lg pull-lg-right">提交</button>
          </form>
        </div>
      </div>
    </div>
  </section>


  {{template "inc/footer.tpl" .}}

</body>

</html>
