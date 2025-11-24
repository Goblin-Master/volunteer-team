export default {} 

import type BaseResp from '@/types/base';

export function mapOk<S, T>(raw: BaseResp<S | null>, to: (s: S) => T): BaseResp<T | null> {
  const data = raw.code === 0 && raw.data ? to(raw.data) : null;
  return { code: raw.code, message: raw.message, data };
}