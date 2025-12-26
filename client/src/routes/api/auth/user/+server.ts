import { PUBLIC_AWS_INVOKE_URL } from '$env/static/public';
import { error, isHttpError, type RequestEvent } from '@sveltejs/kit';

export async function GET(event: RequestEvent) {
	try {
		const invokeUrl = PUBLIC_AWS_INVOKE_URL;
		const cookies = event.request.headers.get('Cookie');

		const res = await fetch(`${invokeUrl}/user`, {
			method: 'GET',
			headers: cookies === null ? undefined : { Cookie: cookies }
		});

		if (!res.ok) {
			throw error(res.status, 'user not found.');
		}

		const data = await res.json();

		if (data.error) {
			throw error(404, 'user not found.');
		}
		if (!data.email || !data.display_name) {
			throw error(404, 'user not found.');
		}

		return new Response(JSON.stringify({ display_name: data.display_name, email: data.email }), {
			status: 200
		});
	} catch (err) {
		if (isHttpError(err)) {
			throw err;
		}

		throw error(500, 'Internal Server Error');
	}
}
