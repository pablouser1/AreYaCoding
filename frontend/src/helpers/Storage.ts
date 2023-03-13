export default class Storage {
  static PREFIX = "areyacoding"

  static getToken(): string {
    return Storage.get("token");
  }

  static getUsername(): string {
    return Storage.get("username");
  }

  static get(key: string): string {
    const val = localStorage.getItem(Storage.getKey(key));
    if (val !== null) {
      return val;
    }

    return "";
  }

  static set(item: string, val: string): void {
    localStorage.setItem(Storage.getKey(item), val);
  }

  private static getKey(item: string): string {
    return `${Storage.PREFIX}-${item}`;
  }
}