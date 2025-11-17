/** @type {import('next').NextConfig} */
const nextConfig = {
reactStrictMode: true,
images: { unoptimized: true }, // для простых деплоев без image-оптимизатора
experimental: { typedRoutes: true }
};
export default nextConfig;