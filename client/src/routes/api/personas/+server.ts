import { error, isHttpError, type RequestEvent } from '@sveltejs/kit';
import { PUBLIC_AWS_INVOKE_URL } from '$env/static/public';

export async function GET(event: RequestEvent) {
	try {
		const cookies = event.request.headers.get('Cookie');
		const invokeUrl = PUBLIC_AWS_INVOKE_URL;
		const headers: Record<string, string> = {
			'Content-Type': 'application/json'
		};
		if (cookies) {
			headers['Cookie'] = cookies;
		}

		const res = await fetch(`${invokeUrl}/personas`, {
			method: 'GET',
			headers: headers
		});

		if (!res.ok) {
			throw error(res.status);
		}

		const data = await res.json();

		if (data.error) {
			throw error(500, data.error);
		}

		return new Response(JSON.stringify(data), { status: 200 });
	} catch (err) {
		if (isHttpError(err)) {
			throw err;
		}
		console.log(err);
		throw error(500, 'internal server error');
	}
}
