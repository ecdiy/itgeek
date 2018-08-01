import base from '@/libs/frame/base.vue';
import frame from '@/libs/frame/frame.vue';

const r = function (c, n, t) {
    return {
        path: '/p/' + n.replace('_', '/'), name: n, component: c, title: t
    }
}

export const routers = [
    {
        path: '/', component: base,
        children: [
            {path: '/', name: 'home', component: () => import('@/base/home.vue'), title: '首页'},

            r(() => import('@/user/login.vue'), 'user_login', '登录'),
            r(() => import('@/user/register.vue'), 'user_register', '注册'),
            r(() => import('@/user/forget.vue'), 'user_forget', '忘记密码'),

            r(() => import('@/user/member.vue'), 'member/:id', ''),

            r(() => import('@/base/403.vue'), '403', '403'),
            r(() => import('@/base/404.vue'), '404', '404'),
            r(() => import('@/base/500.vue'), '500', '500'),

            r(() => import('@/base/home.vue'), 'topic/list,:id', ''),
            r(() => import('@/topic/detail.vue'), 'topic/detail,:id', ''),


            r(() => import('@/down/index.vue'), 'down'),
            r(() => import('@/notes/hot.vue'), 'note'),


            r(() => import('@/down/item.vue'), 'down/item,:s'),
        ]
    },
    {
        auth: true, path: '/', component: base,
        children: [

            r(() => import('@/topic/newMarkdown.vue'), 'newMarkdown', '创作新主题'),
            r(() => import('@/topic/newHTML.vue'), 'newHTML', '创作新主题'),

            r(() => import('@/notes/list.vue'), 'notes_list', '笔记'),
            r(() => import('@/notes/newHtml.vue'), 'notes_newHTML', '创建笔记'),
            r(() => import('@/notes/newMarkdown.vue'), 'notes_newMarkdown', '创建笔记'),
            r(() => import('@/notes/category.vue'), 'notes_category', '笔记分类管理'),
            r(() => import('@/notes/detail.vue'), 'notes/item/:id', ''),


            r(() => import('@/down/my.vue'), 'down_my'),
            r(() => import('@/down/upload.vue'), 'down_upload'),
        ]
    },
    {
        auth: true, path: '/', component: frame,
        children: [
            r(() => import('@/my/balance'), 'my/balance,:id', '账户余额'),
            r(() => import('@/my/setting.vue'), 'my_setting', '设置'),
            r(() => import('@/my/LoginAward'), 'my_LoginAward', '领取今日登录奖励'),
            r(() => import('@/user/topPlayer'), 'user_topPlayer', '社区消费排行榜'),
            r(() => import('@/my/msg.vue'), 'my_msg', '消息'),

            r(() => import('@/topic/append.vue'), 'topic/append,:id', '增加附言'),

        ]
    },
    {
        path: '/', component: frame,
        children: [
            r(() => import('@/user/dau.vue'), 'user_dau', '今日活跃会员'),
            r(() => import('@/my/fav.vue'), 'my_fav', '我的收藏'),
            r(() => import('@/my/follow.vue'), 'my_follow', '特别关注'),

            r(() => import('@/topic/referer.vue'), 'topic_referer', '访问统计'),
        ]
    }

];
