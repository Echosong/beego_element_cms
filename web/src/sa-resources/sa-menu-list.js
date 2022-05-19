/* eslint-disable */

import sa from '../static/sa';
import router from '../router';

var RouteConfig = router;

/* 菜单集合 */
export default new Promise((suc, fail) => {
	sa.get("/permission/listByUser/1").then(res => {
		let returnMenu = res.data.map(item => {
			RouteConfig.forEach(r => {
				if (item.perms == 'home') {
					item.hide_close = true; // 隐藏关闭键 
					item.is_rend = true;
				}
				if (r.perms == item.perms) {
					item.view = r.view;
					if (item.type > 1) {
						item.is_show = item.show;
					}

				}
			})
			return item;
		})
		suc(returnMenu);
	})
})


