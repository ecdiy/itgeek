const webpack = require('webpack');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const ExtractTextPlugin = require('extract-text-webpack-plugin');
// const CopyWebpackPlugin = require('copy-webpack-plugin');
const cleanWebpackPlugin = require('clean-webpack-plugin');
const merge = require('webpack-merge');
const webpackBaseConfig = require('./m.webpack.base.config.js');

var TransformModulesPlugin = require('webpack-transform-modules-plugin')
const path = require('path');

var jobName = process.env.npm_lifecycle_event
var idx = jobName.indexOf('-')
jobName = jobName.substr(idx + 1)

module.exports = merge(webpackBaseConfig, {
    output: {
        publicPath: '/h5dist/',  // 修改 https://iv...admin 这部分为你的服务器域名
        filename: '[name].[hash].js',
        chunkFilename: '[name].[hash].chunk.js',
        path: path.resolve(__dirname, '../dist/' + jobName + '/h5dist')
    },
    plugins: [ new TransformModulesPlugin(),
        new cleanWebpackPlugin(['dist/' + jobName + '/*'], {
            root: path.resolve(__dirname, '../')
        }),
        new ExtractTextPlugin({
            filename: '[name].[hash].css',
            allChunks: true
        }),
        new webpack.optimize.CommonsChunkPlugin({
            // name: 'vendors',
            // filename: 'vendors.[hash].js'
            name: ['vender-exten', 'vender-base'],
            minChunks: Infinity
        }),
        new webpack.DefinePlugin({
            'process.env': {
                NODE_ENV: '"production"'
            }
        }),
        new webpack.optimize.UglifyJsPlugin({
            compress: {
                warnings: false
            }
        }),

        // new CopyWebpackPlugin([
        //     //
        //     //
        //     // {
        //     //     from: jobName + '/components/tinymce'
        //     // }
        // ], {
        //     // ignore: [
        //     //     'text-editor.vue'
        //     // ]
        // }),
        new HtmlWebpackPlugin({
            filename: '../index.html',
            template: '!!ejs-loader!./' + jobName + '/libs/index.ejs',
            inject: false
        })
    ]
});

