import { Axios } from "axios";
import ApiRes from "./models/ApiRes";
import Room from "./models/Room";

export default class Api {
  private client = new Axios({
    baseURL: '/api/v1'
  });

  private username = "";
  private token = "";

  async login(username: string): Promise<string> {
    const res = await this.client.post('/auth/login', null, {
      params: {username}
    });
    return JSON.parse(res.data);
  }

  async getRooms(): Promise<ApiRes<Room[]>> {
    const res = await this.client.get('/rooms');
    return JSON.parse(res.data);
  }

  async createRoom(roomName: string): Promise<ApiRes<Room>> {
    const res = await this.client.post('/rooms', null, {
      headers: {
        Authorization: this.getTokenHeader()
      },
      params: {roomName}
    })

    return JSON.parse(res.data);
  }

  // -- HELPERS -- //
  isLoggedIn(): boolean {
    return this.token !== "";
  }

  getUsername(): string {
    return this.username;
  }

  setAuth(token: string, username: string): void {
    this.token = token;
    this.username = username;
  }

  private getTokenHeader(): string {
    return `Bearer ${this.token}`
  }
}
