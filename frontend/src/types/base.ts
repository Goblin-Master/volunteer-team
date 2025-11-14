export default interface BaseResp<T = any> {
  code: number;
  message: string;
  data: T;
}
