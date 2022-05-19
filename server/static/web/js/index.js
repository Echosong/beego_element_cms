$(function(){
	$(".menu").click(function(){
		if($(this).hasClass("menu-close")){
			$(this).removeClass("menu-close");
			$(".menu_list").slideUp();
			}else{
				$(this).addClass("menu-close");
				$(".menu_list").slideDown();
				}
		});
	$(".foot_tab div span").click(function(){
		$(this).parent().siblings("div").removeClass("active");
		if(!$(this).parent().hasClass("active")){
			$(this).parent().addClass("active");
		}else{
			$(this).parent().removeClass("active");
		}
	})
	$(document).click(function(e){
		var target=$(e.target);
		if(target.closest(".foot_tab div span").length==0&&target.closest(".foot_tab div ul").length==0){
			$(".foot_tab div").removeClass("active");
		}
	})
})
