"use client";

import { useState } from 'react';
import Link from 'next/link';

export default function SignIn() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  const handleSubmit = async (e: { preventDefault: () => void; }) => {
    e.preventDefault();
    try {
      const res = await fetch('/api/signin', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password }),
      });
      if (res.ok) {
        alert('Sign in successful');
      } else {
        const { message } = await res.json();
        setError(message || 'Sign in failed');
      }
    } catch (err) {
      console.error(err);
      setError('Sign in failed');
    }
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-900 text-white">
      <div className="bg-gray-800 p-8 rounded shadow-md w-full max-w-md">
        <h2 className="text-2xl font-bold mb-4">Sign In</h2>
        {error && <p className="text-red-500 mb-4">{error}</p>}
        <form onSubmit={handleSubmit} >
          <div className="mb-4">
            <label className="block text-gray-400">Email</label>
            <input
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              className="w-full px-3 py-2 border rounded text-black"
            />
          </div>
          <div className="mb-6">
            <label className="block text-gray-400">Password</label>
            <input
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="w-full px-3 py-2 border rounded text-black"
            />
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
