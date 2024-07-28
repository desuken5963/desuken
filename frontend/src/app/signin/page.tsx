"use client";

import Link from 'next/link';

export default function SignIn() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-900 text-white">
      <div className="bg-gray-800 p-8 rounded shadow-md w-full max-w-md">
        <h2 className="text-2xl font-bold mb-4">Sign In</h2>
        <form>
          <div className="mb-4">
            <label className="block text-gray-700">Email</label>
            <input type="email" className="w-full px-3 py-2 border rounded" />
          </div>
          <div className="mb-6">
            <label className="block text-gray-700">Password</label>
            <input type="password" className="w-full px-3 py-2 border rounded" />
          </div>
          <button type="submit" className="w-full bg-blue-500 text-white py-2 rounded">Sign In</button>
        </form>
        <p className="mt-4 text-center">
          Don't have an account? <Link href="/signup"><span className="text-blue-500 cursor-pointer">Sign Up</span></Link>
        </p>
      </div>
    </div>
  );
}
