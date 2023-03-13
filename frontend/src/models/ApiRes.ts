export default interface ApiRes<T> {
  status: number;
  success: boolean;
  msg: string;
  data: T;
}
