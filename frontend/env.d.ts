/// <reference types="vite/client" />
// env.d.ts (或者 shims-vue.d.ts)

declare module '*.vue' {
    import type { DefineComponent } from 'vue'
    const component: DefineComponent<{}, {}, any>
    export default component
}