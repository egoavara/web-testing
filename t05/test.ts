type IsObject<T> = T extends object ? true : false;
type LitteralSupportTypes =
  | undefined
  | boolean
  | string
  | null
  | number
  | bigint;

type TypeFlatten<T> = {
  [K in innerTypeFlatten4<T, keyof T>]: innerTypeFlattenValue<T, K>;
};

type innerTypeFlatten4<T, K extends keyof T> = K extends LitteralSupportTypes
  ? IsObject<T[K]> extends true
    ? `${K}.${innerTypeFlatten3<T[K], keyof T[K]>}`
    : `${K}`
  : never;

type innerTypeFlatten3<T, K extends keyof T> = K extends LitteralSupportTypes
  ? IsObject<T[K]> extends true
    ? `${K}.${innerTypeFlatten2<T[K], keyof T[K]>}`
    : `${K}`
  : never;

type innerTypeFlatten2<T, K extends keyof T> = K extends LitteralSupportTypes
  ? IsObject<T[K]> extends true
    ? `${K}.${innerTypeFlatten1<T[K], keyof T[K]>}`
    : `${K}`
  : never;

type innerTypeFlatten1<T, K extends keyof T> = K extends LitteralSupportTypes
  ? IsObject<T[K]> extends true
    ? never
    : `${K}`
  : never;

type innerTypeFlattenValue<
  T,
  K extends string
> = K extends `${infer PRE}.${infer POST}`
  ? PRE extends keyof T
    ? innerTypeFlattenValue<T[PRE], POST>
    : never
  : K extends keyof T
  ? T[K]
  : never;

interface A {
  a: string;
  b: {
    aa: string;
    bb: number;
  };
  c: {
    aa: {
      aaa: string;
      bbb: {
        aaaa: string;
        bbbb: {
          aaaaa: string;
        };
      };
    };
    bb: {
      aaa: string;
    };
  };
}

function flatten<T extends object>(target: T): TypeFlatten<T> {
  return {} as TypeFlatten<T>;
}
const result = flatten({
  a: "key 'a'",
  b: {
    aa: "key 'b.aa'",
    bb: "key 'b.bb'",
  },
  c: {
    aa: {
      aaa: "key 'c.aa.aaa'",
      bbb: "key 'c.aa.bbb'",
      n : 1
    },
  },
});
result[""];
