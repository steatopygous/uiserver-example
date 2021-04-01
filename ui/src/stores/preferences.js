export const preferences = {
    async setWindowSize(width, height) {
        const body = { width, height };

        const response = await fetch(`/api/preferences`, {
            method: 'PATCH',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(body)
        });

        if (!response.ok) {
            console.warn('setWindowSize() - received error response');
        }
    }
}
