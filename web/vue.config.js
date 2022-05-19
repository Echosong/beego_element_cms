module.exports = {
	runtimeCompiler: true,
	devServer:{
		port: 8081,
		proxy: {
			"/api":{
				target: "http://localhost:8080/api",
				changeOrigin: true,
				pathRewrite:{
					'^/api':''
				}
			},
			'/upload':{
				target: "http://localhost:8080/upload",
				changeOrigin: true,
				pathRewrite: {
					'^/upload':''
				}
			}
			
		}
	},
	lintOnSave:false
}