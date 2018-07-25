import Main from '@/base/Main.vue';

const r = function (c, n, t) {
    return {
        path: '/p/' + n.replace('_', '/'), name: n, component: c, title: t
    }
}

export const routers = [

    r(() => import('@/topic/detail.vue'), 'topic/detail,:id', ''),

    r(() => import('@/base/home.vue'), 'topic/list,:id', ''),
    
    {
    path: '/', component: Main,
    children: [
        {path: '/', name: 'home', component: () => import('@/base/home.vue'), title: '首页'},

        r(() => import('@/user/login.vue'), 'user_login', '登录'),
        r(() => import('@/user/register.vue'), 'user_register', '注册'),
        r(() => import('@/user/forget.vue'), 'user_forget', '忘记密码'),

        r(() => import('@/user/member.vue'), 'member/:id', ''),

        r(() => import('@/base/home.vue'), '', '首页'),

    ]
}, {
    path: '/', component: Main,
    children: [
        r(() => import('@/my/setting.vue'), 'my_setting', '设置'),

        // r(() => import('@/notes/note.vue'), 'note_:id', '笔记'),

        // r(() => import('@/notes/list.vue'), 'notes_list', '笔记'),
        // r(() => import('@/notes/new.vue'), 'notes_new', '创建笔记'),
        // r(() => import('@/notes/category.vue'), 'notes_category', '笔记分类管理'),
        // r(() => import('@/notes/detail.vue'), 'notes/item/:id', ''),
    ]
}];
