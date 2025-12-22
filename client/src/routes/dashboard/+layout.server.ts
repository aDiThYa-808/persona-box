import { error, redirect } from '@sveltejs/kit';
import { PUBLIC_BASE_URL, PUBLIC_ENVIRONMENT } from '$env/static/public';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ locals, fetch }) => {
	if (!locals.user) {
		throw redirect(302, '/login');
	}
	try {
		const res = await fetch(`api/auth/user`);

		if (!res.ok) {
			throw error(res.status, 'user not found');
		}

		const data = await res.json();

		return {
			name: data.display_name,
			email: data.email
		};
	} catch (err) {
		throw redirect(302, '/login');
	}
};
