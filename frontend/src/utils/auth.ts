import * as jwt_decode from 'jwt-decode';

interface DecodedToken {
  username: string;
}

export const getUsernameFromToken = (token: string) => {
    const decodedToken = jwt_decode.jwtDecode<DecodedToken>(token);
    return decodedToken.username;
};
