export function useUser() {
  console.log(document.cookie);
  return { id: document.cookie.match(/userId=(?<id>[^;]+);?$/).groups.id };
}
