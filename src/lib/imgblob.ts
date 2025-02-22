export function makeBlobURL(ascii: string): string {
  const binary = atob(ascii)
  const buffer = new Uint8Array(binary.length)
  for (let i = 0; i < binary.length; i++) {
    buffer[i] = binary.charCodeAt(i)
  }
  const blob = new Blob([buffer], { type: 'image/jpeg' })
  return URL.createObjectURL(blob)
}
