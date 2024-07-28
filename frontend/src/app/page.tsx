import Link from 'next/link';

export default function Home() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-900 text-white">
      <div className="text-center">
        <h1 className="text-4xl font-bold mb-4">Welcome to Our Site</h1>
        <p className="mb-4">Please sign in or sign up to continue.</p>
        <div className="space-x-4">
          <Link href="/signin">
            <span className="px-4 py-2 bg-blue-500 text-white rounded cursor-pointer">Sign In</span>
          </Link>
          <Link href="/signup">
            <span className="px-4 py-2 bg-green-500 text-white rounded cursor-pointer">Sign Up</span>
          </Link>
        </div>
      </div>
    </div>
  );
}
