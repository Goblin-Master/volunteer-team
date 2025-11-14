// 登录成功后，Data 字段的类型
export interface LoginResp {
  token: string;
  username: string;
  avatar: string;
}

export interface LoginModel {
  token: string;
  username: string;
  avatar: string;
}

export const toLoginModel = (x: LoginResp): LoginModel => ({
  token: x.token,
  username: x.username,
  avatar: x.avatar,
});

export const toLoginResp = (x: LoginModel): LoginResp => ({
  token: x.token,
  username: x.username,
  avatar: x.avatar,
});
