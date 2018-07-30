const webpack = require('webpack');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const ExtractTextPlugin = require('extract-text-webpack-plugin');
const CopyWebpackPlugin = require('copy-webpack-plugin');
const cleanWebpackPlugin = require('clean-webpack-plugin');
const merge = require('webpack-merge');
const webpackBaseConfig = require('./webpack.base.config.js');

const path = require('path');

var jobName = process.env.npm_lifecycle_event
var idx = jobName.indexOf('-')
jobName = jobName.substr(idx + 1)

module.exports = merge(webpackBaseConfig, {
    output: {
        publicPath: '/admin/dist/',  // 修改  这部分为你的服务器域名
        filename: '[name].[hash].js',
        chunkFilename: '[name].[hash].chunk.js',
        path: path.resolve(__dirname, '../dist/' + jobName + '/dist')
    },
    plugins: [
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

        new CopyWebpackPlugin([

            // {
            //     from: 'itgeek.top/styles/fonts',
            //     to: 'fonts'
            // },
            {
                from: 'components/tinymce'
            }
        ], {
            ignore: [
                'text-editor.vue'
            ]
        }),
        new HtmlWebpackPlugin({
            filename: '../index.html',
            template: '!!ejs-loader!./' + jobName + '/libs/index.ejs',
            inject: false
        })
    ]
});