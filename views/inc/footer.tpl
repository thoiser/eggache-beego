<script src="/static/js/jquery-1.11.3.min.js">
</script>
<script src="/static/js/tether.min.js"></script>
<script src="/static/js/bootstrap.min.js"></script>
<script>
  getUrl = function(name) {
    var str = name.toLowerCase();
    if (location.href.toLowerCase().indexOf(str) != -1) {
      return true;
    } else {
      return false;
    }
  }
  if (getUrl('photo')) {
    $('#leftul li').removeClass('active').eq(0).addClass('active')
  };
  if (getUrl('article')) {
    $('#leftul li').removeClass('active').eq(1).addClass('active')
  };
  if (getUrl('video')) {
    $('#leftul li').removeClass('active').eq(2).addClass('active')
  };
</script>
