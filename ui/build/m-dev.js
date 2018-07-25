const webpack = require('webpack');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const ExtractTextPlugin = require('extract-text-webpack-plugin');
const CopyWebpackPlugin = require('copy-webpack-plugin');
const merge = require('webpack-merge');
const webpackBaseConfig = require('./m.webpack.base.config.js');

const package = require('../package.json');


module.exports = merge(webpackBaseConfig, {
    devtool: '#source-map',
    output: {
        publicPath: '/dist/',
        filename: '[name].js',
        chunkFilename: '[name].chunk.js'
    },
    plugins: [
        new ExtractTextPlugin({
            filename: '[name].css',
            allChunks: true
        }),
        new webpack.optimize.CommonsChunkPlugin({
            name: ['vender-exten', 'vender-base'],
            minChunks: Infinity
        }),
        new HtmlWebpackPlugin({
            title: 'itgeek.top ',
            filename: '../index.html',
            inject: false
        }),
        // new CopyWebpackPlugin([
        //     {
        //         from: 'admin.kushow.me/components/tinymce'
        //     }
        // ], {
        //     ignore: [
        //         'text-editor.vue'
        //     ]
        // })
    ]
});