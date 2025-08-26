export default interface BaseResp<T = any> {
    code: number; // 注意：Go 后端返回的是小写
    msg: string;
    data: T;
}