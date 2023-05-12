config_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd ${config_dir}

${config_dir}/build_linux.sh

scp log_server_demo2_linux ecs-user@182.92.99.61:
scp .env ecs-user@182.92.99.61:

ssh ecs-user@182.92.99.61 '
    if [ ! -d 'test2' ]; then mkdir test2; fi;
    mv log_server_demo2_linux test2
    mv .env test2
'