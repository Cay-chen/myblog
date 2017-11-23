{{define "header"}}
<div class="headtop"></div>
    <div class="contenttop">
      <div class="logo f_l">Cay-Chen个人博客</div>
      <div class="search f_r">
        <form action="/listpic" method="post" name="searchform" id="searchform">
          <input name="keyboard" id="keyboard" class="input_text" value="请输入关键字" style="color: rgb(153, 153, 153);" onfocus="if(value=='请输入关键字'){this.style.color='#000';value=''}" onblur="if(value==''){this.style.color='#999';value='请输入关键字'}" type="text">
          <input name="page" value="1" type="hidden">
          <input name="Submit" class="input_submit" value="搜索" type="submit">
        </form>
      </div>
      <div class="blank"></div>
      <nav>
        <div  class="navigation">
          <ul class="menu">
            <li><a href="/">网站首页</a></li>
             <!-- <ul>
                <li><a href="about.html">个人简介</a></li>
                <li><a href="listpic.html">个人相册</a></li>
              </ul>-->
            <li><a href="/listpic?classify=101&page=1">慢生活</a></li>
            <li><a href="/listpic?page=1">学习笔记</a>
              <ul>
                <li><a href="/listpic?classify=103&page=1">Golang</a></li>
                <li><a href="/listpic?classify=105&page=1">MySql</a></li>
                <li><a href="/listpic?classify=106&page=1">Java</a></li>
                <li><a href="/listpic?classify=102&page=1">Android</a></li>
                <li><a href="/listpic?classify=104&page=1">PHP</a></li>
                <li><a href="/listpic?classify=108&page=1">JS/JQ</a></li>
                <li><a href="/listpic?classify=107&page=1">HTML/CSS</a></li>
                <li><a href="/listpic?classify=109&page=1">jQuery</a></li>
              </ul>
            </li>
            <li><a href="/aboutus" target="_blank">我的爱情</a>
            <li><a href="/3dphoto" target="_blank">3D相册</a> </li>
            <li><a href="#">给我留言</a> </li>
          </ul>
        </div>
      </nav>
      <SCRIPT type=text/javascript>
	// Navigation Menu
	$(function() {
		$(".menu ul").css({display: "none"}); // Opera Fix
		$(".menu li").hover(function(){
			$(this).find('ul:first').css({visibility: "visible",display: "none"}).slideDown("normal");
		},function(){
			$(this).find('ul:first').css({visibility: "hidden"});
		});
	});
</SCRIPT>
    </div>
{{end}}