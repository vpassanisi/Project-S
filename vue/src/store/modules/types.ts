export interface Sub {
  _id: string;
  name: string;
  description: string;
  founder: string;
  banner: string;
  createdAt: number;
  colorPrimary: string;
  colorPrimaryLight: string;
  colorPrimaryDark: string;
}

export interface User {
  name: string;
}

export interface Post {
  _id: string;
  title: string;
  body: string;
  user: User;
  points: number;
  sub: string;
  createdAt: number;
}

export interface Comment {
  _id: string;
  body: string;
  points: number;
  user: string;
  post: string;
  parent: string;
  createdAt: number;
}
