import { error, redirect } from '@sveltejs/kit';
import type { User } from '$lib/types/user';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ locals, fetch, params }) => {
	if (!locals.user) {
		throw redirect(302, '/login');
	}
	try {
		const res = await fetch(`/api/auth/user`,{method:"GET"});

		if (!res.ok) {
			throw error(res.status, 'user not found');
		}

		const userData = await res.json();

		const pageData = {
			name: userData.display_name,
			email: userData.email,
			personaid : params.personaid,
			chatid: params.chatid
		};

		return pageData;
	} catch (err) {
		throw redirect(302, '/login');
	}
};
