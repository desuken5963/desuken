"use client";

import Link from 'next/link';

export default function SignUp() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-900 text-white">
      <div className="bg-gray-800 p-8 rounded shadow-md w-full max-w-md">
        <h2 className="text-2xl font-bold mb-4">Sign Up</h2>
        <form>
          <div className="mb-4">
            <label className="block text-gray-700">Email</label>
            <input type="email" className="w-full px-3 py-2 border rounded" />
          </div>
          <div className="mb-4">
            <label className="block text-gray-700">Password</label>
            <input type="password" className="w-full px-3 py-2 border rounded" />
          </div>
          <div className="mb-6">
            <label className="block text-gray-700">Confirm Password</label>
            <input type="password" className="w-full px-3 py-2 border rounded" />
          </div>
          <button type="submit" className="w-full bg-green-500 text-white py-2 rounded">Sign Up</button>
        </form>
        <p className="mt-4 text-center">
          Already have an account? <Link href="/signin"><span className="text-blue-500 cursor-pointer">Sign In</span></Link>
        </p>
      </div>
    </div>
  );
}
