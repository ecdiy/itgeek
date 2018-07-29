import gk from '@/libs/frame/gk.vue';

const r = function (c, n, p) {
    return {
        path: p ? p : '/p/' + n.replace('_', '/'), name: n, component: c
    }
};

export const routers = [

    r(() => import('@/base/login.vue'), 'login'),
    {
        path: '/', component: gk,
        children: [
            r(() => import('@/base/home.vue'), 'home', '/'),
            r(() => import('@/gkadmin/index.vue'), 'gkadmin_index,:id'),
            r(() => import('@/gkadmin/category.vue'), 'gkadmin_category,:id'),
            r(() => import('@/gkadmin/scoreRule.vue'), 'gkadmin_scoreRule,:id'),

            r(() => import('@/gkadmin/user.vue'), 'gkadmin_user,:id'),

            r(() => import('@/zx/pageIndex.vue'), 'gkadmin_pageIndex,:id'),
            r(() => import('@/zx/pageTopicDetail.vue'), 'gkadmin_pageTopicDetail,:id'),
            r(() => import('@/gkadmin/topic.vue'), 'gkadmin_topic,:id'),
        ]
    }

];
