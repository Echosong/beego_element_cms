var h,s=$(".cp_tab a").length;

$(function(){

	getCss();
	$(window).resize(function(){
		getCss();
scrolls();
	})
if(window.performance.navigation.type==1){
scrolls();
}else{
loadCur();
}
	$(".cp_tab a").click(function(){
		var index=$(this).index();
		$("html").animate({
			scrollTop: $(".cp_list li").eq(index).find(".div").offset().top
		}, 500,"swing",function(){

			$(".cp_tab a").eq(index).addClass("current").siblings().removeClass("current");
		});	
	})
	$(window).scroll(function(){
		scrolls();
	})
})
function getCss(){
	if($(window).width()>1100){
		h=$(".cp_tab").outerHeight()+$(".header").outerHeight()
	}else{
		h=$(".cp_tab").outerHeight()+$(".nav").outerHeight()
	}
	$(".div").css({top:-h+"px"});         
}
function scrolls(){
if(!$("html").is(":animated")){
	for(var i=0; i<s; i++){
		if(i==0){
			if($("html").scrollTop()<$(".cp_list li").eq(i).find(".div").offset().top+0.8*$(".cp_list li").eq(i).outerHeight()){
				$(".cp_tab a").eq(i).addClass("current").siblings().removeClass("current");
return false
			}
		}else if(i==s-1){
			if($("html").scrollTop()>=$(".cp_list li").eq(i).find(".div").offset().top-0.2*$(".cp_list li").eq(i).outerHeight()){
				$(".cp_tab a").eq(i).addClass("current").siblings().removeClass("current");
return false
			}
		}else{
			if($("html").scrollTop()>=$(".cp_list li").eq(i).find(".div").offset().top-0.2*$(".cp_list li").eq(i).outerHeight()&&$("html").scrollTop()<$(".cp_list li").eq(i).find(".div").offset().top+0.8*$(".cp_list li").eq(i).outerHeight()){
				$(".cp_tab a").eq(i).addClass("current").siblings().removeClass("current");
return false
			}
		}
	}

}
}
function loadCur(){
	var id=location.hash;
if(id==""){
	$(".cp_tab a").eq(0).addClass("current");
}else{
	var index=$(id).parent().index();
	$(".cp_tab a").eq(index).addClass("current");
}

}