<template>
  <div v-if="shouldShow" class="notice-container">
    <div class="notice-circle" @click="toggleModal">
      <span class="circle-text">!</span>
    </div>
    <div v-if="is_modal_open" class="notice-modal-overlay" @click.self="toggleModal">
      <div class="notice-modal-content">
        <button class="close-btn" @click="toggleModal">&times;</button>
        <div class="notice-text-content">
          <h4>电脑报修须知:</h4>
          <ol>
            <li>网络问题: 可先到学校网络中心进行报修;
              <p>松山湖校区网络中心电话: 22861378 </p>
              <p>大朗、莞城、校本部网络中心电话: 22680756/13377692407/13377692259</p>
            </li>
            <li>师友志愿服务团队因受平台限制，目前尚未支持硬件故障报修，若同学电脑硬件出现故障，请联系厂家或自行解决，造成不便，敬请谅解。</li>
            <li>报修时，请尽量详细地描述电脑故障！</li>
            <li>成功提交报修申请后，请耐心等待工作人员联系！</li>
            <li>服务时间一般为工作日，我们会在1-2个工作日内与你取得联系，如有不便，敬请原谅。</li>
            <li>您的报修信息，我们将保密处理！</li>
            <li>如有疑问或紧急情况，请联系:
              <p>松山湖校区: </p>
              <p>陆晓彤 长号: 14778189977 微信号: hanyu12072022</p>
              <p>刘奕豪 长号: 13714721480 微信号: sa582560</p>
            </li>
          </ol>
          <h4>Tip:</h4>
          <ul>
            <li>莞城校区报修系统尚未开放，敬请谅解。</li>
            <li>电话联系仅用于紧急联系，对学生只提供咨询（课后时间），敬请谅解。</li>
            <li>请确认您填写的联系方式，若有错误工作人员可能无法与您取得联系。</li>
            <li>师友是莞工唯一志愿为全校师生提供志愿的免费电脑维修服务的院级组织，虽然身为学生我们在时间和能力上有所许，不能及的地方，但我们承诺尽我们所能为你解决问题，所以请大家尊重我们的劳动，因为你的支持与鼓励是我们继续志愿服务的最大动力！</li>
            <li>若检测发现您的电脑是由于主板或其他硬件造成物理性故障，由于平台问题我们无法提供维修服务。但会指引您或帮助您进行联系厂商进行维修。</li>
            <li>师友提供的硅脂比较一般，散热性能一般，如果同学或者老师介意的话建议自带硅脂，若出现因硅脂质量问题造成的维修概不负责。</li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import { useRoute } from 'vue-router';

// State
const is_modal_open = ref(false);

// Router
const route = useRoute();
const login_paths = ['/login', '/register', '/resetPassword']; // 根据你的路由路径进行修改

// Computed
const shouldShow = computed(() => {
  return !login_paths.includes(route.path);
});

// Methods
const toggleModal = () => {
  is_modal_open.value = !is_modal_open.value;
};
</script>

<style scoped>
/* 悬浮小圆圈样式 */
.notice-container {
  position: fixed;
  bottom: 20px;
  right: 20px;
  z-index: 1000;
}

.notice-circle {
  width: 50px;
  height: 50px;
  background-color: #409EFF; /* 按钮颜色，可自定义 */
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.2);
  transition: transform 0.2s;
}

.notice-circle:hover {
  transform: scale(1.1);
}

.circle-text {
  color: white;
  font-size: 24px;
  font-weight: bold;
}

/* 弹窗样式 */
.notice-modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.6);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1001;
}

.notice-modal-content {
  background: white;
  padding: 30px;
  border-radius: 10px;
  max-width: 80%;
  max-height: 80%;
  overflow-y: auto;
  position: relative;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.3);
}

.close-btn {
  position: absolute;
  top: 10px;
  right: 10px;
  background: none;
  border: none;
  font-size: 24px;
  cursor: pointer;
}

.notice-text-content {
  line-height: 1.6;
}
</style>
