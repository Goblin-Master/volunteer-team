// src/api/file.ts
import req from './axios';
import type BaseResp  from '@/types/base.ts'; // 引入你的 BaseResp 类型

/**
 * 上传图片文件的 API 接口。
 * @param file 要上传的图片文件对象。
 * @returns 返回一个 Promise，成功时其 data 字段为图片在服务器上的 URL。
 */
export const updateAvatar = (file: File): Promise<BaseResp<string>> => {
  // FormData 用于构建 multipart/form-data 类型的请求体，
  // 这是上传文件时常用的方式。
  const formData = new FormData();
  // 'file' 是后端接口用来接收文件的字段名，请根据你的后端接口文档进行修改。
  formData.append('file', file);

  return req({
    url: '/api/user/updateAvatar', // 这是假设的图片上传接口地址，请替换为你的实际地址。
    method: 'post',
    data: formData,
    headers: {
      // 必须设置 'Content-Type': 'multipart/form-data'，
      // 但浏览器会自动为 FormData 设置，所以可以省略。
      // 不过这里为了清晰，我们还是列出来。
      'Content-Type': 'multipart/form-data',
    },
  });
};