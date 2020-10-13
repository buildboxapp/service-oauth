function bblogin(redirect, vendors, theme) {
	let width = 600, height = 450, left = (screen.width - width)/2, top = (screen.height - height)/2;
	return window.open('https://buildbox.app/auth/login/?vendors='+encodeURIComponent(vendors||'')+
		'&theme='+encodeURIComponent(theme||'')+
		'&redirect='+encodeURIComponent(redirect||document.location.href),
		'bboauth', `left=`+left+`,top=`+top+`,width=`+width+`,height=`+height+`,
        menubar=no,toolbar=no,location=no,status=yes`);
}

if (jQuery) {
	(function( $ ){
		$.fn.bblogin = function(redirect, vendors, theme) {
			return this.on('click', function(e){
				e.preventDefault();
				bblogin(redirect, vendors, theme);
			});
		};
	})(jQuery);
}